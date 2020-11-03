import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_blog/blogapi.dart';
import 'package:mango_folio/bodies/banner.dart';
import 'package:mango_folio/bodies/colour.dart';
import 'package:mango_folio/bodies/contact.dart';
import 'package:mango_folio/bodies/content.dart';
import 'package:mango_folio/bodies/information.dart';
import 'package:mango_folio/bodies/section.dart';
import 'package:mango_folio/bodies/simpleblock.dart';
import 'package:mango_folio/contentapi.dart';
import 'package:mango_ui/keys.dart';

void main() {
  querySelector('#btnAdd').onClick.listen(onAddClick);
  document.body.onClick.matches('.deleter').listen(onDeleteClick);
}

void onAddClick(MouseEvent e) async {
  final data = new Content(
      "none",
      "empty",
      new Key("0`0"),
      "en",
      "your@email.com",
      new Banner(new Key("0`0"), new Key("0`0"), "Empty", "Subtitle"),
      new Section("Section A Header", "Subtitle", "", new Key("0`0")),
      new Section("Section B Header", "Subtitle", "", new Key("0`0")),
      new Information("Info", "Information text", new List<SimpleBlock>()),
      new Colour(
          new RGB(199, 95, 37, "#c75f25"),
          new RGB(199, 95, 37, "#c75f25"),
          new RGB(199, 95, 37, "#c75f25"),
          new RGB(199, 95, 37, "#c75f25"),
          new RGB(199, 95, 37, "#c75f25"),
          new RGB(199, 95, 37, "#c75f25")),
      new List<Contact>(),
      "Description for client",
      "0-000");
  final req = await createContent(data);

  if (req.status == 200) {
    final key = jsonDecode(req.response);

    final redir = "/content/${key}";
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
