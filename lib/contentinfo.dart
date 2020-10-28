import 'dart:html';

import 'package:FOLIO.APP/bodies/blockitem.dart';
import 'package:mango_cms/bodies/information.dart';
import 'package:mango_cms/bodies/simpleblock.dart';

class ContentInfoForm {
  TextInputElement txtInfoHeader;
  TextInputElement txtInfoText;

  ContentInfoForm() {
    txtInfoHeader = querySelector("#txtInfoHeader");
    txtInfoText = querySelector("#txtContentInfoText");
  }

  String get heading {
    return txtInfoHeader.value;
  }

  String get text {
    return txtInfoText.value;
  }

  List<SimpleBlock> get blocks {
    return findItems();
  }

  List<SimpleBlock> findItems() {
    var isLoaded = false;
    var result = new List<SimpleBlock>();
    var indx = 0;

    do {
      var item =
          new BlockItem("#txtContentText${indx}", "#txtContentIcon${indx}");

      isLoaded = item.loaded;
      if (isLoaded) {
        result.add(item.toDTO());
      }

      indx++;
    } while (isLoaded);

    return result;
  }

  Information toDTO() {
    return new Information(heading, text, blocks);
  }
}
