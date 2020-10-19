import 'banner.dart';
import 'colour.dart';
import 'information.dart';
import 'section.dart';

class Content {
  final String profile;
  final String language;
  final Banner banner;
  final Section sectionA;
  final Section sectionB;
  final Information info;
  final Colour colour;

  Content(this.profile, this.language, this.banner, this.sectionA,
      this.sectionB, this.info, this.colour);

  Map<String, dynamic> toJson() {
    return {
      "Profile": profile,
      "Language": language,
      "Banner": banner,
      "SectionA": sectionA,
      "SectionB": sectionB,
      "Info": info,
      "Colour": colour,
    };
  }
}
