import 'dart:html';

import 'package:FOLIO.APP/bodies/addressitem.dart';
import 'package:mango_entity/bodies/address.dart';

class AddressesForm {
  void onAddClick(MouseEvent e) {
    addItem();
  }

  List<Address> get items {
    return findItems();
  }

  List<Address> findItems() {
    var isLoaded = false;
    var result = new List<Address>();
    var indx = 0;

    do {
      var item = new AddressItem(
          "#numAddrStreetNo${indx}",
          "#txtAddrStreet${indx}",
          "#txtAddrUnitNo${indx}",
          "#txtAddrEstate${indx}",
          "#txtAddrSuburb${indx}",
          "#txtAddrCity${indx}",
          "#txtAddrProvince${indx}",
          "#txtAddrPostalCode${indx}",
          "#txtAddrCoordinates${indx}",
          "#chkIsDeliver${indx}");

      isLoaded = item.loaded;
      if (isLoaded) {
        result.add(item.toDTO());
      }

      indx++;
    } while (isLoaded);

    return result;
  }

  void addItem() {
    //var schema = buildElement(items.length);
    //form.children.add(schema);
  }
}
