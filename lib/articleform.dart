import 'dart:convert';
import 'dart:html';

import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_blog/blogapi.dart';
import 'package:mango_blog/bodies/article.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

class ArticleForm extends FormState {
  Key _objKey;
  TextInputElement _title;
  TextInputElement _intro;
  SelectElement _categories;
  DivElement _content;
  FileUploadInputElement _headImage;
  HiddenInputElement _author;
  CheckboxInputElement _public;

  ArticleForm(
      String idElem,
      Key objKey,
      String titleElem,
      String introElem,
      String categoriesElem,
      String contentElem,
      String imageElem,
      String authorElem,
      String publicElem,
      String submitBtn)
      : super(idElem, submitBtn) {
    _objKey = objKey;
    _title = querySelector(titleElem);
    _intro = querySelector(introElem);
    _categories = querySelector(categoriesElem);
    _content = querySelector(contentElem);
    _headImage = querySelector(imageElem);
    _author = querySelector(authorElem);
    _public = querySelector(publicElem);

    querySelector(submitBtn).onClick.listen(onSubmitClick);
    //querySelector(previewBtn).onClick.listen(onPreviewClick);
    //querySelector(publishBtn).onClick.listen(onPublishClick);

    _headImage.onChange.listen(uploadFile);

    //Editor events
    querySelectorAll('#editCtrls')
        .onClick
        .matches('span.button')
        .listen(onEditorCtrlClick);
  }

  String get title {
    return _title.value;
  }

  String get intro {
    return _intro.value;
  }

  String get category {
    return _categories.value;
  }

  String get content {
    return _content.innerHtml;
  }

  Key get imageKey {
    return new Key(_headImage.dataset['id']);
  }

  String get writtenby {
    return _author.value;
  }

  bool get public {
    return _public.checked;
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final obj = new Article(
          title, intro, category, imageKey, content, writtenby, public);

      HttpRequest req;
      if (_objKey.toJson() != "0`0") {
        req = await updateArticle(_objKey, obj);
      } else {
        req = await createArticle(obj);
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

  void onPreviewClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final obj = new Article(
          title, intro, category, imageKey, content, writtenby, public);
      final req = await updateArticle(_objKey, obj);

      if (req.status == 200) {
        window.open("/blog/articles/${_objKey}", '_blank');
      }
    }
  }

  void onPublishClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      _public.checked = true;

      final obj = new Article(
          title, intro, category, imageKey, content, writtenby, public);
      final req = await updateArticle(_objKey, obj);
      var result = jsonDecode(req.response);

      if (req.status == 200) {
        window.alert(result['Data']);
      } else {
        print(result['Error']);
      }
    }
  }

  void onEditorCtrlClick(MouseEvent e) {
    final ctrl = e.matchingTarget;

    if (ctrl is SpanElement) {
      final role = ctrl.dataset['role'];

      switch (role) {
        case 'h1':
        case 'h2':
        case 'p':
          document.execCommand('formatBlock', false, role);
          break;
        default:
          document.execCommand(role, false, null);
          break;
      }
    }
  }
}
