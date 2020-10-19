import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/content.dart';

Future<HttpRequest> createContent(Content obj) async {
  var apiroute = getEndpoint("cms");
  var url = "${apiroute}/content";
  var data = jsonEncode(obj.toJson());
  return invokeService("POST", url, data);
}

Future<HttpRequest> updateContent(Key key, Content obj) async {
  var apiroute = getEndpoint("cms");
  var url = "${apiroute}/content/${key.toJson()}";
  var data = jsonEncode(obj.toJson());
  return invokeService("PUT", url, data);
}
