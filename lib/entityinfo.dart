import 'dart:html';

class EntityInfoForm {
  TextInputElement txtName;
  SelectElement ddlProfileKey;
  TextInputElement txtIdentification;

  EntityInfoForm() {
    txtName = querySelector("#txtName");
    ddlProfileKey = querySelector("#ddlProfileKey");
    txtIdentification = querySelector("#txtIdentification");

    //btnaddressadd
  }

  String get name {
    return txtName.value;
  }

  String get profileKey {
    return ddlProfileKey.value;
  }

  String get identification {
    return txtIdentification.value;
  }
}
