package handlers

import (
	"testing"
)

func TestGetVideoById(t *testing.T) {
	cases := []struct {
		in   string
		want VideoInfo
		err  string
	}{
		{"-1", VideoInfo{}, "no video found"},
		{"d290f1ee-6c54-4b01-90e6-d701748f0851", VideoInfo{
			Id:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
			Name:      "Black Retrospective Woman",
			Duration:  15,
			Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
			Url:       "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
		}, ""},
		{"hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345", VideoInfo{
			Id:        "hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
			Name:      "Танцор",
			Duration:  91,
			Thumbnail: "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
			Url:       "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/index.mp4",
		}, ""},
		{"sldjfl34-dfgj-523k-jk34-5jk3j45klj34", VideoInfo{
			Id:        "sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
			Name:      "Go Rally TEASER-HD",
			Duration:  41,
			Thumbnail: "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
			Url:       "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/index.mp4",
		}, ""},
	}

	for _, c := range cases {
		got, err := GetVideoById(c.in)
		if (got != c.want) || (err != nil && err.Error() != c.err) {
			t.Errorf("GetVideoById(%v) == %q, wanted %q, error must be %q, not %q", c.in, got, c.want, c.err, err)
		}
	}
}
