package pkg

import (
	"UnitConverter/model"
)

func ConvertLength(post model.Request) model.Response {
	switch post.ConvertFrom {
	case "millimeter":
		switch post.ConvertTo {
		case "centimeter":
			return model.Response{Result: post.Meaning / 10}
		case "meter":
			return model.Response{Result: post.Meaning / 10 / 100}
		case "kilometer":
			return model.Response{Result: post.Meaning / 10 / 100 / 1000}
		}

	case "centimeter":
		switch post.ConvertTo {
		case "millimeter":
			return model.Response{Result: post.Meaning * 10}
		case "meter":
			return model.Response{Result: post.Meaning / 100}
		case "kilometer":
			return model.Response{Result: post.Meaning / 100 / 1000}
		}

	case "meter":
		switch post.ConvertTo {
		case "millimeter":
			return model.Response{Result: post.Meaning * 10 * 100}
		case "centimeter":
			return model.Response{Result: post.Meaning * 100}
		case "kilometer":
			return model.Response{Result: post.Meaning / 1000}
		}

	case "kilometer":
		switch post.ConvertTo {
		case "millimeter":
			return model.Response{Result: post.Meaning * 10 * 100 * 1000}
		case "centimeter":
			return model.Response{Result: post.Meaning * 100 * 1000}
		case "meter":
			return model.Response{Result: post.Meaning * 1000}
		}
	}
	return model.Response{}
}
