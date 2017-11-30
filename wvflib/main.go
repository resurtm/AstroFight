package wvflib

type Vertex struct {
	X, Y, Z, W float64
}

type TexCoord struct {
	U, V float64
}

type Normal struct {
	X, Y, Z, W float64
}

type Object struct {
	Name string
}

type MaterialLibrary struct {
	Name      string
	Materials map[string]Material
}

type MaterialColor struct {
	R, G, B float64
}

type MaterialSpectral struct {
	// file.rfl
	Name string
	Factor float64
}

type MaterialCIEXYZ struct {
	X, Y, Z float64
}

type Material struct {
	Name string

	// Ka
	Ambient struct {
		// Ka r g b
		Color MaterialColor
		// Ka spectral file.rfl factor
		Spectral MaterialSpectral
		// Ka xyz x y z
		CIEXYZ MaterialCIEXYZ
	}

	// Kd
	Diffuse struct {
		// Kd r g b
		Color MaterialColor
		// Kd spectral file.rfl factor
		Spectral MaterialSpectral
		// Kd xyz x y z
		CIEXYZ MaterialCIEXYZ
	}

	// Ks
	Specular struct {
		// Ks r g b
		Color MaterialColor
		// Ks spectral file.rfl factor
		Spectral MaterialSpectral
		// Ks xyz x y z
		CIEXYZ MaterialCIEXYZ
	}

	// Tf
	TransmissionFilter struct {
		// Tf r g b
		Color MaterialColor
		// Tf spectral file.rfl factor
		Spectral MaterialSpectral
		// Tf xyz x y z
		CIEXYZ MaterialCIEXYZ
	}


}
