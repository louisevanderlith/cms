import 'dart:html';

import 'package:mango_secure/bodies/contact.dart';
import 'package:mango_ui/trustvalidator.dart';

import 'bodies/contactitem.dart';

class ContactsForm {
  DivElement form;

  ContactsForm() {
    form = querySelector("#dvContacts");
    querySelector("#btnAddContact").onClick.listen(onAddClick);
  }

  void onAddClick(MouseEvent e) {
    addItem();
  }

  List<Contact> get items {
    return findItems();
  }

  List<Contact> findItems() {
    var isLoaded = false;
    var result = new List<Contact>();
    var indx = 0;

    do {
      var item = new ContactItem("#txtContactName${indx}",
          "#txtContactIcon${indx}", "#txtContactValue${indx}");

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
                    <a href="#" data-liID="liContact${index}" class="card-header-icon" aria-label="more options">
                        <span class="icon">
                            <i class="fas fas-close" aria-hidden="true"></i>
                        </span>
                    </a>
                </header>
                <div class="card-content">
                    <div class="content">
                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtContactName${index}">Name</label>
                                <input class="input" type="text" min-length="3" id="txtContactName${index}"
                                       required/>
                                <p class="help is-danger"></p>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtContactValue${index}">Value</label>
                                <input class="input" type="text" min-length="3" id="txtContactValue${index}"
                                       required/>
                                <p class="help is-danger"></p>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtContactIcon${index}">Icon</label>
                                <input class="input" type="text" id="txtContactIcon${index}" required/>
                                <p class="help is-danger"></p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        ''';

    return new Element.html(schema, validator: new TrustedNodeValidator());
  }
}
