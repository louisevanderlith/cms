class Colour {
  final RGB primary;
  final RGB secondary;
  final RGB tertiary;
  final RGB shadows;
  final RGB accent;
  final RGB background;

  Colour(this.primary, this.secondary, this.tertiary, this.shadows, this.accent,
      this.background);

  Map<String, dynamic> toJson() {
    return {
      "Primary": primary,
      "Secondary": secondary,
      "Tertiary": tertiary,
      "Shadows": shadows,
      "Accent": accent,
      "Background": background
    };
  }
}

class RGB {
  final num red;
  final num green;
  final num blue;
  final String hex;

  RGB(this.red, this.green, this.blue, this.hex);

  Map<String, dynamic> toJson() {
    return {
      "Red": red,
      "Green": green,
      "Blue": blue,
      "Hex": hex,
    };
  }
}
