import 'package:mango_cms/bodies/simpleblock.dart';

class Information {
  final String heading;
  final String text;
  final List<SimpleBlock> blocks;

  Information(this.heading, this.text, this.blocks);

  Map<String, dynamic> toJson() {
    return {
      "Heading": heading,
      "Text": text,
      "Blocks": blocks
    };
  }
}
