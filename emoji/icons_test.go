package emoji

import "testing"

func TestIcons(t *testing.T) {
	if CHECK_ICON == "" {
		t.Fatal("CHECK_ICON is empty")
	}
	if TOOLS_ICON != WRENCH_ICON+HAMMER_ICON {
		t.Fatalf("unexpected TOOLS_ICON: %s", TOOLS_ICON)
	}
	if APPAREL_ICON != DRESS_ICON+HIGH_HEELED_SHOES_ICON+TSHIRT_ICON {
		t.Fatalf("unexpected APPAREL_ICON: %s", APPAREL_ICON)
	}
	if MAIN_MENU_ICON != HOME_ICON {
		t.Fatalf("unexpected MAIN_MENU_ICON: %s", MAIN_MENU_ICON)
	}
}
