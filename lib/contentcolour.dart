import 'dart:html';

import 'package:mango_cms/bodies/colour.dart';

class ContentColourForm {
  InputElement colPrimary;
  InputElement colSecondary;
  InputElement colTertiary;
  InputElement colShadows;
  InputElement colAccent;
  InputElement colBackground;

  ContentColourForm() {
    colPrimary = querySelector("#colPrimary");
    colSecondary = querySelector("#colSecondary");
    colTertiary = querySelector("#colTertiary");
    colShadows = querySelector("#colShadows");
    colAccent = querySelector("#colAccent");
    colBackground = querySelector("#colBackground");
  }

  RGB get primary {
    return hexToRGB(colPrimary.value);
  }

  RGB get secondary {
    return hexToRGB(colSecondary.value);
  }

  RGB get tertiary {
    return hexToRGB(colTertiary.value);
  }

  RGB get shadows {
    return hexToRGB(colShadows.value);
  }

  RGB get accent {
    return hexToRGB(colAccent.value);
  }

  RGB get background {
    return hexToRGB(colBackground.value);
  }

  RGB hexToRGB(String hex) {
    final rgx = new RegExp(r"^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$",
        caseSensitive: false, multiLine: false);

    final val = rgx.firstMatch(hex);

    final r = int.parse(val.group(1), radix: 16);
    final g = int.parse(val.group(2), radix: 16);
    final b = int.parse(val.group(3), radix: 16);

    return new RGB(r, g, b, hex);
  }

  Colour toDTO() {
    return new Colour(
        primary, secondary, tertiary, shadows, accent, background);
  }
}
