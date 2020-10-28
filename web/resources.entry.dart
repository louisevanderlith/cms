import 'dart:html';

import 'package:FOLIO.APP/resourceform.dart';
import 'package:mango_ui/keys.dart';

ButtonElement btnCreate;

void main() {
  btnCreate = querySelector("#btnCreate");
  btnCreate.onClick.listen(showCreate);

  querySelector("#btnCloseModal").onClick.listen(closeModal);
}

void showCreate(MouseEvent e) {
  new ResourceForm(new Key("0%600"));

  btnCreate.disabled = true;
  querySelector("#theModal").classes.add('is-active');
}

void closeModal(MouseEvent e) {
  btnCreate.disabled = false;
  querySelector("#theModal").classes.remove('is-active');
}
