import 'dart:html';

import 'package:mango_entity/bodies/address.dart';

class AddressItem {
  bool _loaded;

  NumberInputElement numStreetNo;
  TextInputElement txtStreet;
  TextInputElement txtUnitNo;
  TextInputElement txtEstateName;
  TextInputElement txtSuburb;
  TextInputElement txtCity;
  TextInputElement txtProvince;
  TextInputElement txtPostalCode;
  TextInputElement txtCoordinates;
  CheckboxInputElement chkIsDelivery;

  AddressItem(
      String streetNoId,
      String streetId,
      String unitNoId,
      String estateId,
      String suburbId,
      String cityId,
      String provinceId,
      String postalCodeId,
      String coordinatesId,
      String isDeliveryId) {
    numStreetNo = querySelector(streetNoId);
    txtStreet = querySelector(streetId);
    txtUnitNo = querySelector(unitNoId);
    txtEstateName = querySelector(estateId);
    txtSuburb = querySelector(suburbId);
    txtCity = querySelector(cityId);
    txtProvince = querySelector(provinceId);
    txtPostalCode = querySelector(postalCodeId);
    txtCoordinates = querySelector(coordinatesId);
    chkIsDelivery = querySelector(isDeliveryId);
  }

  bool get loaded {
    return _loaded;
  }

  num get streetNo {
    return numStreetNo.valueAsNumber;
  }

  String get street {
    return txtStreet.value;
  }

  String get unitNo {
    return txtUnitNo.value;
  }

  String get estateName {
    return txtEstateName.value;
  }

  String get suburb {
    return txtSuburb.value;
  }

  String get city {
    return txtCity.value;
  }

  String get province {
    return txtProvince.value;
  }

  String get postalCode {
    return txtPostalCode.value;
  }

  String get coordinates {
    return txtCoordinates.value;
  }

  bool get isDelivery {
    return chkIsDelivery.checked;
  }

  Address toDTO() {
    return new Address(streetNo, street, unitNo, estateName, suburb, city,
        province, postalCode, coordinates, isDelivery);
  }
}
