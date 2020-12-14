package models

type PackType string

const (
	PT_Base PackType = "BASE"
	PT_SD            = "SD"
	PT_HD            = "HD"
	PT_4K            = "4K"
)

type Pack struct {
	Type  PackType
	Price float64
}
