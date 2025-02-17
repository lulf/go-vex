/*
Copyright 2023 The OpenVEX Authors
SPDX-License-Identifier: Apache-2.0
*/

package csaf

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	doc, err := Open("testdata/csaf.json")
	require.NoError(t, err)
	require.NotNil(t, doc)
	require.Equal(t, "Example VEX Document", doc.Document.Title)
	require.Equal(t, "CSAFPID-0001", doc.FirstProductName())

	// Vulnerabilities
	require.Len(t, doc.Vulnerabilities, 1)
	require.Equal(t, doc.Vulnerabilities[0].CVE, "CVE-2009-4487")
	require.Len(t, doc.Vulnerabilities[0].ProductStatus, 1)
	require.Len(t, doc.Vulnerabilities[0].ProductStatus["known_not_affected"], 1)
	require.Equal(t, doc.Vulnerabilities[0].ProductStatus["known_not_affected"][0], "CSAFPID-0001")
}

func TestFindFirstProduct(t *testing.T) {
	doc, err := Open("testdata/csaf.json")
	require.NoError(t, err)
	require.NotNil(t, doc)

	prod := doc.ProductTree.FindFirstProduct()
	require.Equal(t, prod, "CSAFPID-0001")
}

func TestFindByHelper(t *testing.T) {
	doc, err := Open("testdata/csaf.json")
	require.NoError(t, err)
	require.NotNil(t, doc)

	prod := doc.ProductTree.FindProductIdentifier("purl", "pkg:maven/@1.3.4")
	require.NotNil(t, prod)
	require.Equal(t, prod.ID, "CSAFPID-0001")
}

func TestListProducts(t *testing.T) {
	doc, err := Open("testdata/csaf.json")
	require.NoError(t, err)
	require.NotNil(t, doc)
	prods := doc.ProductTree.Branches[0].Branches[0].Branches[0].ListProducts()
	require.Len(t, prods, 1)
	require.Equal(t, prods[0].IdentificationHelper["purl"], "pkg:golang/github.com/go-homedir@v1.1.0")

	allProds := doc.ProductTree.Branches[0].ListProducts()
	require.NotNil(t, allProds)
	require.Len(t, allProds, 2)
}
