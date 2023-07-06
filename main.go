package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-logr/stdr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

//go:embed all:static
var staticFiles embed.FS

var (
	tmplt                 *template.Template
	lastFetchedTime       int64
	cachedIngresses       []IngressesData
	cachedNamespaces      []string
	annotationsIgnoreList = []string{
		"kubectl.kubernetes.io/last-applied-configuration",
		"kubernetes.io/ingress.class",
		"kubernetes.io/tls-acme",
		"nginx.ingress.kubernetes.io/auth-signin",
		"nginx.ingress.kubernetes.io/auth-url",
		"cert-manager.io/cluster-issuer",
	}
)

type TemplateData struct {
	Ingresses  []IngressesData
	Namespaces []string
}

type IngressesData struct {
	Name        string
	Link        string
	Namespace   string
	Annotations map[string]string
}

func getPageData() (TemplateData, error) {
	currentTime := time.Now().Unix()

	if currentTime-lastFetchedTime > 15 {
		fmt.Printf("Cache MISS: Attempting to fetching ingress data from cluster at %d\n", currentTime)
		lastFetchedTime = currentTime
		cachedNamespaces = []string{}

		// Fetch ingresses from the cluster with one K8s API call
		ctx := context.Background()
		ctrl.SetLogger(
			stdr.New(
				log.New(os.Stdout, "", log.Lshortfile),
			),
		)
		config := ctrl.GetConfigOrDie()
		clientset := kubernetes.NewForConfigOrDie(config)
		ingresses, err := clientset.NetworkingV1().Ingresses("").List(ctx, metav1.ListOptions{})
		if err != nil {
			return TemplateData{}, err
		}

		// Compile ingresses data
		ingressData := []IngressesData{}
		for _, ingress := range ingresses.Items {
			// Filter annotations
			filteredAnnotations := map[string]string{}
			for key, value := range ingress.Annotations {
				ignored := false
				for _, ignoreKey := range annotationsIgnoreList {
					if key == ignoreKey {
						ignored = true
						break
					}
				}
				if !ignored {
					filteredAnnotations[key] = value
				}
			}
			adder := IngressesData{
				Name:        ingress.Name,
				Link:        "https://" + ingress.Spec.Rules[0].Host,
				Namespace:   ingress.Namespace,
				Annotations: filteredAnnotations,
			}

			// Compile namespace list
			present := false
			for _, namespace := range cachedNamespaces {
				if namespace == ingress.Namespace {
					present = true
					break
				}
			}
			if !present {
				cachedNamespaces = append(cachedNamespaces, ingress.Namespace)
			}
			ingressData = append(ingressData, adder)
		}
		cachedIngresses = ingressData
	} else {
		fmt.Printf("Cache HIT: Using cache from %d (%ds ago)\n", lastFetchedTime, currentTime-lastFetchedTime)
	}

	data := TemplateData{
		Ingresses:  cachedIngresses,
		Namespaces: cachedNamespaces,
	}

	return data, nil
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		tmplt, _ = template.ParseFiles("site.html")
		data, err := getPageData()
		if err != nil {
			panic(
				fmt.Sprintf("Failed to get page data: %s", err),
			)
		}

		err = tmplt.Execute(writer, data)

		if err != nil {
			return
		}
	}
}

func runServer() {
	// Set up static filesystem
	var staticFS = fs.FS(staticFiles)
	static, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.FS(static))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlePage)

	err = http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalln("Server failed to start:", err)
	}
}

func main() {
	lastFetchedTime = 0
	cachedIngresses = []IngressesData{}
	runServer()
}
