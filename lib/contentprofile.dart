import 'dart:html';

import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_ui/keys.dart';

class ContentProfileForm {
  TextInputElement txtProfileRealm;
  TextInputElement txtProfileClient;
  TextInputElement txtProfileLanguage;
  EmailInputElement txtProfileEmail;
  FileUploadInputElement uplProfileLogo;
  TextAreaElement txtProfileDescription;
  TextInputElement txtProfileGTag;

  ContentProfileForm() {
    txtProfileRealm = querySelector("#txtProfileRealm");
    txtProfileClient = querySelector("#txtProfileClient");
    txtProfileLanguage = querySelector("#txtProfileLanguage");
    txtProfileEmail = querySelector("#txtProfileEmail");
    uplProfileLogo = querySelector("#uplProfileLogoImg");
    txtProfileDescription = querySelector("#txtProfileDescription");
    txtProfileGTag = querySelector("#txtProfileGTag");

    uplProfileLogo.onChange.listen(uploadFile);
  }

  String get realm {
    return txtProfileRealm.value;
  }

  String get client {
    return txtProfileClient.value;
  }

  String get language {
    return txtProfileLanguage.value;
  }

  String get email {
    return txtProfileEmail.value;
  }

  Key get logo {
    return new Key(uplProfileLogo.dataset['id']);
  }

  String get description {
    return txtProfileDescription.value;
  }

  String get gtag {
    return txtProfileGTag.value;
  }
}
