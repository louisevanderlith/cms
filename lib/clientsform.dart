import 'dart:html';

import 'package:mango_secure/bodies/client.dart';
import 'package:mango_ui/trustvalidator.dart';

import 'bodies/clientitem.dart';

class ClientsForm {
  DivElement form;

  ClientsForm() {
    form = querySelector("#dvClients");
    querySelector("#btnAddClient").onClick.listen(onAddClick);
  }

  void onAddClick(MouseEvent e) {
    addItem();
  }

  List<Client> get items {
    return findItems();
  }

  List<Client> findItems() {
    var isLoaded = false;
    var result = new List<Client>();
    var indx = 0;

    do {
      var item = new ClientItem(
          "#txtClientName${indx}",
          "#txtClientSecret${indx}",
          "#txtClientUrl${indx}",
          "#lstResources${indx}",
          "#chkClientCodes${indx}",
          "#chkClientTerms${indx}");

      isLoaded = item.loaded;
      if (isLoaded) {
        result.add(item.toDTO());
      }

      indx++;
    } while (isLoaded);

    return result;
  }

  void addItem() {
    var schema = buildElement(items.length);
    form.children.add(schema);
  }

  //returns HTML for this Item
  Element buildElement(int index) {
    var schema = '''
    <div class="card">
                <header class="card-header">
                    <a href="#" data-liID="liClient${index}" class="card-header-icon" aria-label="more options"> </a>
                </header>
                <div class="card-content">
                    <div class="content">
                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtClientName${index}">Name</label>
                                <input class="input" type="text" min-length="3" id="txtClientName${index}"
                                       required/>
                                <p class="help is-danger"></p>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtClientSecret${index}">Secret</label>
                                <input class="input" type="password" id="txtClientSecret${index}"
                                       required/>
                                <p class="help is-danger"></p>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtClientUrl${index}">Url</label>
                                <input class="input" type="url" id="txtClientUrl${index}" required/>
                                <p class="help is-danger"></p>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                                <label class="label" for="chkClientTerms${index}">Terms Enabled</label>
                                <input class="input" type="checkbox" id="chkClientTerms${index}" required
                                       checked="false"/>
                                <p class="help is-danger"></p>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                                <label class="label" for="chkClientCodes${index}">Codes Enabled</label>
                                <input class="input" type="text" id="chkClientCodes${index}" required
                                       checked="false"/>
                                <p class="help is-danger"></p>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                            <ul id="lstResources${index}"></ul>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        ''';

    return new Element.html(schema, validator: new TrustedNodeValidator());
  }
}
