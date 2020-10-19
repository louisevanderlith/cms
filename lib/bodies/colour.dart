class Colour {
  final num red;
  final num green;
  final num blue;
  final String hex;

  Colour(this.red, this.green, this.blue, this.hex);

  Map<String, dynamic> toJson() {
    return {
      "Red": red,
      "Green": green,
      "Blue": blue,
      "Hex": hex,
    };
  }
}
