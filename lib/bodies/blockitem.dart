import 'dart:html';

import 'package:mango_folio/bodies/simpleblock.dart';

class BlockItem {
  TextInputElement txtIcon;
  TextInputElement txtText;

  bool _loaded;

  BlockItem(String textId, String iconId) {
    txtText = querySelector(textId);
    txtIcon = querySelector(iconId);

    _loaded = txtIcon != null && txtText != null;
  }

  bool get loaded {
    return _loaded;
  }

  String get icon {
    return txtIcon.value;
  }

  String get text {
    return txtText.value;
  }

  SimpleBlock toDTO() {
    return new SimpleBlock(icon, text);
  }
}
