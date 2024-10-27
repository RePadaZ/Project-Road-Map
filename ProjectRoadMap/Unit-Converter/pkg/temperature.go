package pkg

import (
	"UnitConverter/model"
)

func ConvertTemp(post model.Request) model.Response {
	switch post.ConvertFrom {
	case "celsius":
		switch post.ConvertTo {
		case "fahrenheit":
			return model.Response{Result: post.Meaning*1.8 + 32}
		case "kelvin":
			return model.Response{Result: post.Meaning + 273.15}
		}

	case "fahrenheit":
		switch post.ConvertTo {
		case "celsius":
			return model.Response{Result: (post.Meaning - 32) / 1.8}
		case "kelvin":
			return model.Response{Result: (post.Meaning + 459.67) * 5 / 9}
		}

	case "Kelvin":
		switch post.ConvertTo {
		case "celsius":
			return model.Response{Result: post.Meaning - 273.15}
		case "fahrenheit":
			return model.Response{Result: post.Meaning*9/5 - 459.67}
		}

	}
	return model.Response{}
}
