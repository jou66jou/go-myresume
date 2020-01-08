package main

import (
	"github.com/phpdave11/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: gopdf.Rect{W: 595.28, H: 841.89}, //595.28, 841.89 = A4
		Protection: gopdf.PDFProtectionConfig{
			UseProtection: true,
			Permissions:   gopdf.PermissionsPrint | gopdf.PermissionsCopy | gopdf.PermissionsModify,
			OwnerPass:     []byte("123"),
			UserPass:      []byte("123"),
		},
	})

	pdf.AddPage()
	pdf.Image("./123.png", 10, 60, &gopdf.Rect{W: 595.28, H: 780}) //print image
	pdf.WritePdf("protect2.pdf")
}
