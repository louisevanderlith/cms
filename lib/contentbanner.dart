import 'dart:html';

import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_folio/bodies/banner.dart';
import 'package:mango_ui/keys.dart';

class ContentBannerForm {
  FileUploadInputElement uplBackground;
  FileUploadInputElement uplImage;
  TextInputElement txtHeading;
  TextInputElement txtSubtitle;

  ContentBannerForm() {
    uplBackground = querySelector("#uplBannerBackgroundImg");
    uplBackground.onChange.listen(uploadFile);

    uplImage = querySelector("#uplBannerImg");
    uplImage.onChange.listen(uploadFile);

    txtHeading = querySelector("#txtBannerHeading");
    txtSubtitle = querySelector("#txtBannerSubtitle");
  }

  Key get background {
    return new Key(uplBackground.dataset['id']);
  }

  Key get image {
    return new Key(uplImage.dataset['id']);
  }

  String get heading {
    return txtHeading.value;
  }

  String get subtitle {
    return txtSubtitle.value;
  }

  Banner toDTO() {
    return new Banner(background, image, heading, subtitle);
  }
}
