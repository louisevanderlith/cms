import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_secure/bodies/resource.dart';
import 'package:mango_secure/resourceapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

class ResourceForm extends FormState {
  Key _objKey;
  TextInputElement txtName;
  TextInputElement txtDisplayName;
  PasswordInputElement txtSecret;
  TextInputElement txtNewNeed;
  UListElement lstNeeds;

  ResourceForm(Key objKey) : super("#frmResource", "#btnSave") {
    _objKey = objKey;

    txtName = querySelector("#txtName");
    txtDisplayName = querySelector("#txtDisplayName");
    txtSecret = querySelector("#txtSecret");
    txtNewNeed = querySelector("#txtNewNeed");
    lstNeeds = querySelector("#lstNeeds");

    querySelector("#btnSave").onClick.listen(onSubmitClick);
    querySelector("#btnAddNeed").onClick.listen(onNeedAddClick);
  }

  String get name {
    return txtName.value;
  }

  String get displayname {
    return txtDisplayName.value;
  }

  String get secret {
    return txtSecret.value;
  }

  List<String> get needs {
    return lstNeeds.children.map((e) => e.text).toList();
  }

  void onNeedAddClick(MouseEvent e) {
    final item = new Element.li();
    item.text = txtNewNeed.value;

    lstNeeds.children.add(item);
    txtNewNeed.value = "";
  }

  void onSubmitClick(MouseEvent e) async {
    if (!isFormValid()) {
      return;
    }

    disableSubmit(true);

    final obj =
        new Resource(this.name, this.displayname, this.secret, this.needs);

    HttpRequest req;
    if (_objKey.toJson() != "0`0") {
      req = await updateResource(_objKey, obj);
    } else {
      req = await createResource(obj);
    }

    if (req.status == 200) {
      final result = jsonDecode(req.response);
      print(result);
      final data = result['Data'];
      final rec = data['Record'];

      new Toast.success(
          title: "Success!", message: data, position: ToastPos.bottomLeft);

      if (rec != null) {
        final key = rec['K'];

        _objKey = key;
      }
    } else {
      new Toast.error(
          title: "Failed!",
          message: req.response,
          position: ToastPos.bottomLeft);
    }
  }
}
