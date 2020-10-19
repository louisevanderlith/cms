import 'package:mango_ui/keys.dart';

class Section {
  final String heading;
  final String text;
  final String imageUrl;
  final Key imageKey;

  Section(this.heading, this.text, this.imageUrl, this.imageKey);

  Map<String, dynamic> toJson() {
    return {
      "Heading": heading,
      "Text": text,
      "ImageUrl": imageUrl,
      "ImageKey": imageKey
    };
  }
}
