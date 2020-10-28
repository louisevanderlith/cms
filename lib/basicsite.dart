import 'dart:html';

import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_secure/bodies/mapitem.dart';
import 'package:mango_ui/keys.dart';

class BasicSiteForm {
  TextInputElement txtTitle;
  TextAreaElement txtDescription;
  FileUploadInputElement uplProfileImg;

  BasicSiteForm() {
    txtTitle = querySelector("#txtTitle");
    txtDescription = querySelector("#txtDescription");
    uplProfileImg = querySelector("#uplProfileImg");
    uplProfileImg.onChange.listen(uploadFile);

    querySelector("#btnAddCode").onClick.listen(addCodeClick);
    querySelector("#btnAddTerms").onClick.listen(addTermClick);
  }

  void addCodeClick(MouseEvent e) {}

  void addTermClick(MouseEvent e) {}

  String get title {
    return txtTitle.value;
  }

  String get description {
    return txtDescription.text;
  }

  Key get imageKey {
    return new Key(uplProfileImg.dataset['id']);
  }

  List<MapItem> get endpoints {
    return getMapItems("#txtEndpointName", "#txtEndpointValue");
  }

  List<MapItem> get codes {
    return getMapItems("#txtCodeName", "#txtCodeValue");
  }

  List<MapItem> get terms {
    return getMapItems("#txtTermName", "#txtTermValue");
  }
}
