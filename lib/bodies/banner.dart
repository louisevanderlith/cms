import 'package:mango_ui/keys.dart';

class Banner {
  final Key background;
  final Key image;
  final String heading;
  final String subtitle;

  Banner(this.background, this.image, this.heading, this.subtitle);

  Map<String, dynamic> toJson() {
    return {
      "Background": background,
      "Image": image,
      "Heading": heading,
      "Subtitle": subtitle,
    };
  }
}
