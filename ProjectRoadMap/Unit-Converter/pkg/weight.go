package pkg

import (
	"UnitConverter/model"
)

func ConvertWeight(post model.Request) model.Response {
	switch post.ConvertFrom {
	case "milligram":
		switch post.ConvertTo {
		case "gram":
			return model.Response{Result: post.Meaning / 1000}
		case "kilogram":
			return model.Response{Result: post.Meaning / 1000 / 1000}
		}

	case "gram":
		switch post.ConvertTo {
		case "milligram":
			return model.Response{Result: post.Meaning * 1000}
		case "kilogram":
			return model.Response{Result: post.Meaning / 1000}
		}

	case "kilogram":
		switch post.ConvertTo {
		case "milligram":
			return model.Response{Result: post.Meaning * 1000}
		case "gram":
			return model.Response{Result: post.Meaning * 1000 * 1000}
		}

	}
	return model.Response{}
}
