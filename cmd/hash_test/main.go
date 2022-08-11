package main

import (
	"fmt"
	"hash"
	"hash/fnv"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	rukpakv1alpha1 "github.com/operator-framework/rukpak/api/v1alpha1"
	"k8s.io/apimachinery/pkg/util/rand"
	// plain "github.com/operator-framework/rukpak/internal/provisioner/plain/types"
)

func DeepHashObject(hasher hash.Hash, objectToWrite interface{}) {
	hasher.Reset()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	printer.Fprintf(hasher, "%#v", objectToWrite)
}

func GenerateTemplateHash(template interface{}) string {
	hasher := fnv.New32a()
	DeepHashObject(hasher, template)
	return rand.SafeEncodeString(fmt.Sprint(hasher.Sum32()))
}

func main() {
	fmt.Printf("go version: %q, GOOS: %q, GOARCH: %q\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	test_template := rukpakv1alpha1.BundleTemplate{
		Spec: rukpakv1alpha1.BundleSpec{
			ProvisionerClassName: "core.rukpak.io/plain",
			Source: rukpakv1alpha1.BundleSource{
				Type: rukpakv1alpha1.SourceTypeImage,
				Image: &rukpakv1alpha1.ImageSource{
					Ref: "testdata/bundles/plain-v0:valid",
				},
			},
		},
	}
	test_pointer := &rukpakv1alpha1.BundleTemplate{
		Spec: rukpakv1alpha1.BundleSpec{
			ProvisionerClassName: "core.rukpak.io/plain",
			Source: rukpakv1alpha1.BundleSource{
				Type: rukpakv1alpha1.SourceTypeImage,
				Image: &rukpakv1alpha1.ImageSource{
					Ref: "testdata/bundles/plain-v0:valid",
				},
			},
		},
	}
	hashed_template := GenerateTemplateHash(test_template)
	hashed_pointer := GenerateTemplateHash(test_pointer)
	fmt.Printf("hash output (template): %+v\n", hashed_template)
	fmt.Printf("hash output (pointer): %+v\n", hashed_pointer)
}
