package engine

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type ui_text struct {
	Content string
	start_x int32
	start_y int32
	end_x   int32
	end_y   int32
}

func (self *ui_text) MouseDown() bool {
	return mouse_state == sdl.BUTTON_LEFT && self.MouseOver()
}

func (self *ui_text) MouseOver() bool {
	return MouseX >= self.start_x &&
		MouseX <= self.end_x &&
		MouseY >= self.start_y &&
		MouseY <= self.end_y
}

func (self *ui_text) Draw() {
	src_rect := sdl.Rect{0, 0, char_width, char_height}
	dst_rect := sdl.Rect{0, self.start_y, scaled_char_width, scaled_char_height}
	for i, c := range self.Content {
		ascii := int32(c) - 32
		if ascii > 95 || i < 0 {
			log.Println("Character", c, "is not in charmap")
			continue
		}
		charmap_x := ascii % cols
		charmap_y := ascii / cols

		src_rect.X = charmap_x * char_width
		src_rect.Y = charmap_y * char_height

		dst_rect.X = self.start_x + int32(i)*scaled_char_width

		renderer.Copy(
			charmap,
			&src_rect,
			&dst_rect,
		)
	}
}

func calcStartPos(el_size int32, rel_pos float64, screen_size float64, align Alignment) int32 {
	switch align {
	case UI_ALIGN_START:
		return int32(rel_pos * screen_size)
	case UI_ALIGN_CENTER:
		return int32(rel_pos*screen_size) - el_size/2
	case UI_ALIGN_END:
		return int32(rel_pos*screen_size) - el_size
	default:
		log.Println("Invalid alignment", align)
		return 0
	}
}

func CreateAlignedText(content string, rel_x float64, rel_y float64, v_align Alignment, h_align Alignment) UIElement {
	text_width := int32(len(content)) * scaled_char_width
	text_height := scaled_char_height

	abs_x := calcStartPos(text_width, rel_x, width_f64, h_align)
	abs_y := calcStartPos(text_height, rel_y, height_f64, v_align)

	text := ui_text{
		Content: content,
		start_x: abs_x,
		start_y: abs_y,
		end_x:   abs_x + text_width,
		end_y:   abs_y + text_height,
	}
	return &text
}