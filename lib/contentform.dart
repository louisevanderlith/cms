import 'dart:html';

import 'package:CMS.APP/contactsform.dart';
import 'package:dart_toast/dart_toast.dart';
import 'package:mango_folio/bodies/content.dart';
import 'package:mango_folio/contentapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

import 'contentbanner.dart';
import 'contentcolour.dart';
import 'contentinfo.dart';
import 'contentprofile.dart';
import 'contentsection.dart';

class ContentForm extends FormState {
  Key objKey;

  ContentProfileForm profileForm;
  ContentBannerForm bannerForm;
  ContentSectionForm sectionAForm;
  ContentSectionForm sectionBForm;
  ContentInfoForm infoForm;
  ContentColourForm colourForm;
  ContactsForm contactsForm;

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
    contactsForm = new ContactsForm();

    querySelector("#btnSubmit").onClick.listen(onSubmitClick);
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final obj = new Content(
          profileForm.realm,
          profileForm.client,
          profileForm.logo,
          profileForm.language,
          profileForm.email,
          bannerForm.toDTO(),
          sectionAForm.toDTO(),
          sectionBForm.toDTO(),
          infoForm.toDTO(),
          colourForm.toDTO(),
          contactsForm.items,
          profileForm.description,
          profileForm.gtag);

      HttpRequest req;
      if (objKey.toJson() != "0`0") {
        req = await updateContent(objKey, obj);
        if (req.status == 200) {
          Toast.success(
              title: "Success!",
              message: req.response,
              position: ToastPos.bottomLeft);
        } else {
          Toast.error(
              title: "Failed!",
              message: req.response,
              position: ToastPos.bottomLeft);
        }
      } else {
        req = await createContent(obj);

        if (req.status == 200) {
          final key = req.response;
          objKey = new Key(key);

          Toast.success(
              title: "Success!",
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
}
