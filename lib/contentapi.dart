import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/content.dart';

Future<HttpRequest> createContent(Content data) async {
  var apiroute = getEndpoint("cms");
  var url = "${apiroute}/content";

  return invokeService("POST", url, data);
}

Future<HttpRequest> updateContent(Key key, Content data) async {
  var apiroute = getEndpoint("cms");
  var url = "${apiroute}/content/${key.toString()}";

  return invokeService("PUT", url, data);
}
