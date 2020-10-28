import 'dart:html';

import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_cms/bodies/section.dart';
import 'package:mango_ui/keys.dart';

class ContentSectionForm {
  TextInputElement txtHeading;
  TextAreaElement txtText;
  TextInputElement txtImageUrl;
  FileUploadInputElement uplImage;

  ContentSectionForm(String headingElem, String textElem, String imageUrlElem,
      String imageElem) {
    txtHeading = querySelector(headingElem);
    txtText = querySelector(textElem);
    txtImageUrl = querySelector(imageUrlElem);
    uplImage = querySelector(imageElem);
    uplImage.onChange.listen(uploadFile);
  }

  String get heading {
    return txtHeading.value;
  }

  String get text {
    return txtText.value;
  }

  String get imageUrl {
    return txtImageUrl.value;
  }

  Key get image {
    return new Key(uplImage.dataset['id']);
  }

  Section toDTO() {
    return new Section(heading, text, imageUrl, image);
  }
}
