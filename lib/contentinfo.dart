import 'dart:html';

import 'package:CMS.APP/bodies/blockitem.dart';
import 'package:mango_folio/bodies/information.dart';
import 'package:mango_folio/bodies/simpleblock.dart';
import 'package:mango_ui/trustvalidator.dart';

class ContentInfoForm {
  TextInputElement txtInfoHeader;
  TextInputElement txtInfoText;
  DivElement lstBlocks;

  ContentInfoForm() {
    txtInfoHeader = querySelector("#txtInfoHeader");
    txtInfoText = querySelector("#txtContentInfoText");
    lstBlocks = querySelector("#lstBlocks");
    querySelector("#btnAddBlock").onClick.listen(onAddClick);
  }

  void onAddClick(MouseEvent e) {
    addItem();
  }

  String get heading {
    return txtInfoHeader.value;
  }

  String get text {
    return txtInfoText.value;
  }

  List<SimpleBlock> get blocks {
    return findItems();
  }

  List<SimpleBlock> findItems() {
    var isLoaded = false;
    var result = new List<SimpleBlock>();
    var indx = 0;

    do {
      var item =
          new BlockItem("#txtContentText${indx}", "#txtContentIcon${indx}");

      isLoaded = item.loaded;
      if (isLoaded) {
        result.add(item.toDTO());
      }

      indx++;
    } while (isLoaded);

    return result;
  }

  Information toDTO() {
    return new Information(heading, text, blocks);
  }

  void addItem() {
    var schema = buildElement(blocks.length);
    lstBlocks.children.add(schema);
  }

  Element buildElement(num index) {
    final schema = '''
      <div class="card">
        <header class="card-header">
    <p class="card-header-title">
    New Block
    </p>
    <a class="card-header-icon" aria-label="more options">
    <span class="icon">
    <i class="fas fa-angle-down" aria-hidden="true"></i>
    </span>
    </a>
    </header>
    <div class="card-content" hidden>
    <div class="content">
    <div class="field">
    <label class="label" for="txtContentIcon${index}">Icon</label>
    <div class="control">
    <input class="input" type="text" id="txtContentIcon${index}" required
    value=""/>
    <p class="help is-danger"></p>
    </div>
    </div>
    <div class="field">
    <label class="label" for="txtContentText${index}">Text</label>
    <div class="control">
    <input class="input" type="text" id="txtContentText${index}" required
    value=""/>
    <p class="help is-danger"></p>
    </div>
    </div>
    </div>
    <footer class="card-footer">
    <a href="#" data-id="${index}" class="card-footer-item Delete">Delete</a>
    </footer>
    </div>
    </div>
    ''';

    return new Element.html(schema, validator: new TrustedNodeValidator());
  }
}
