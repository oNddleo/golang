package abstract

type GUIFactory interface {
	createButton()
	createCheckbox()
}

func GetGUIFactory(system_os string) (GUIFactory, error) {
	if system_os == "win" {
		return &Window{}, nil
	}

	if brand == "mac" {
		return &MacOs{}, nil
	}

	return nil, fmt.Errorf("Wrong system os type passed")
}
