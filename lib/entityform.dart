import 'dart:convert';
import 'dart:html';

import 'package:FOLIO.APP/entityinfo.dart';
import 'package:FOLIO.APP/usesrform.dart';
import 'package:mango_entity/bodies/address.dart';
import 'package:mango_entity/bodies/entity.dart';
import 'package:mango_entity/bodies/user.dart';
import 'package:mango_entity/entityapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

import 'addressform.dart';

class EntityForm extends FormState {
  Key _objKey;

  EntityInfoForm infoForm;
  AddressesForm addressesForm;
  UserForm userForm;

  EntityForm(Key k) : super("#frmEntity", "#btnSubmit") {
    _objKey = k;

    infoForm = new EntityInfoForm();

    querySelector("#btnSubmit").onClick.listen(onSubmitClick);
  }

  List<Address> get address {
    return addressesForm.items;
  }

  User get user {
    return userForm.toDTO();
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final obj = new Entity(infoForm.name, infoForm.profileKey, user,
          infoForm.identification, address);

      HttpRequest req;
      if (_objKey.toJson() != "0`0") {
        req = await updateEntity(_objKey, obj);
      } else {
        req = await createEntity(obj);
      }

      var result = jsonDecode(req.response);

      if (req.status == 200) {
        final data = result['Data'];
        final rec = data['Record'];

        if (rec != null) {
          final key = rec['K'];
          _objKey = key;
        }
      }
    }
  }
}
