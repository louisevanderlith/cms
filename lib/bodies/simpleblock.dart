class SimpleBlock {
  final String icon;
  final String text;

  SimpleBlock(this.icon, this.text);

  Map<String, dynamic> toJson() {
    return {
      "Icon": icon,
      "Text": text,
    };
  }
}
