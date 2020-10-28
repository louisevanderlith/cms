import 'dart:convert';
import 'dart:html';

import 'package:FOLIO.APP/contentinfo.dart';
import 'package:FOLIO.APP/contentprofile.dart';
import 'package:dart_toast/dart_toast.dart';
import 'package:mango_cms/bodies/content.dart';
import 'package:mango_cms/contentapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

import 'contentbanner.dart';
import 'contentcolour.dart';
import 'contentsection.dart';

class ContentForm extends FormState {
  Key objKey;

  ContentProfileForm profileForm;
  ContentBannerForm bannerForm;
  ContentSectionForm sectionAForm;
  ContentSectionForm sectionBForm;
  ContentInfoForm infoForm;
  ContentColourForm colourForm;

  ContentForm(Key k) : super("#frmContent", "#btnSubmit") {
    objKey = k;

    profileForm = new ContentProfileForm();
    bannerForm = new ContentBannerForm();
    sectionAForm = new ContentSectionForm("#txtSectionAHeader",
        "#txtSectionAText", "#txtSectionAImageURL", "#uplSectionAImg");
    sectionBForm = new ContentSectionForm("#txtSectionBHeader",
        "#txtSectionBText", "#txtSectionBImageURL", "#uplSectionBImg");
    infoForm = new ContentInfoForm();
    colourForm = new ContentColourForm();

    querySelector("#btnSubmit").onClick.listen(onSubmitClick);
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final obj = new Content(
          profileForm.name,
          profileForm.language,
          bannerForm.toDTO(),
          sectionAForm.toDTO(),
          sectionBForm.toDTO(),
          infoForm.toDTO(),
          colourForm.toDTO());

      HttpRequest req;
      if (objKey.toJson() != "0`0") {
        req = await updateContent(objKey, obj);
      } else {
        req = await createContent(obj);
      }

      if (req.status == 200) {
        var result = jsonDecode(req.response);
        
        final data = result['Data'];
        final rec = data['Record'];

        if (rec != null) {
          final key = rec['K'];
          objKey = key;
        }

        Toast.success(
            title: "Saved!",
            message: "Content Saved",
            position: ToastPos.bottomLeft);
      } else {
        Toast.error(
            title: "Failed!",
            message: req.response,
            position: ToastPos.bottomLeft);
      }
    }
  }
}
