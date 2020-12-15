import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_blog/blogapi.dart';
import 'package:mango_blog/bodies/article.dart';
import 'package:mango_ui/keys.dart';

void main() {
  querySelector('#btnAdd').onClick.listen(onAddClick);
  document.body.onClick.matches('.deleter').listen(onDeleteClick);
}

void onAddClick(MouseEvent e) async {
  final data = new Article('New Article', 'Short introduction', 'Default',
      new Key('0`0'), 'Content', 'System', false);
  final req = await createArticle(data);

  if (req.status == 200) {
    final key = jsonDecode(req.response);

    final redir = "/articles/${key}";
    window.location.replace(redir);
  } else {
    Toast.error(
        title: "Failed!", message: req.response, position: ToastPos.bottomLeft);
  }
}

void onDeleteClick(MouseEvent e) async {
  final targt = e.matchingTarget;

  if (targt is ButtonElement) {
    final toDelete = targt.dataset["key"];
    final warn = "Are you sure you want to Delete ${toDelete}?";
    if (window.confirm(warn)) {
      final req = await deleteArticle(new Key(toDelete));

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
    }
  }
}
