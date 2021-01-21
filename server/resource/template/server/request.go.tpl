package request

import "Driving-school/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}