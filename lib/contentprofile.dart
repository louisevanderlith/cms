import 'dart:html';

class ContentProfileForm {
  TextInputElement txtProfileName;
  TextInputElement txtProfileLanguage;

  ContentProfileForm() {
    txtProfileName = querySelector("#txtProfileName");
    txtProfileLanguage = querySelector("#txtProfileLanguage");
  }

  String get name {
    return txtProfileName.value;
  }

  String get language {
    return txtProfileLanguage.value;
  }
}
