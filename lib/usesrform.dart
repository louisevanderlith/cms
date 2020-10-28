import 'dart:html';

import 'package:FOLIO.APP/contactsform.dart';
import 'package:mango_entity/bodies/role.dart';
import 'package:mango_entity/bodies/user.dart';
import 'package:mango_secure/bodies/contact.dart';

class UserForm {
  TextInputElement txtName;
  CheckboxInputElement chkVerified;
  EmailInputElement txtEmail;
  PasswordInputElement txtPassword;
  ContactsForm frmContacts;
  UListElement ulResources;
  UListElement ulRoles;

  UserForm(String nameId, String verifiedId, String emailId, String passwordId,
      String resourceId, String rolesId) {
    txtName = querySelector(nameId);
    chkVerified = querySelector(verifiedId);
    txtEmail = querySelector(emailId);
    txtPassword = querySelector(passwordId);
    frmContacts = new ContactsForm();
    ulResources = querySelector(resourceId);
    ulRoles = querySelector(rolesId);
  }

  String get name {
    return txtName.value;
  }

  bool get isVerified {
    return chkVerified.checked;
  }

  String get email {
    return txtEmail.value;
  }

  String get password {
    return txtPassword.value;
  }

  List<Contact> get contacts {
    return frmContacts.items;
  }

  List<String> get resources {
    return ulResources.children.map((e) => e.text);
    new List<String>();
  }

  List<Role> get roles {
    return new List<Role>();
  }

  User toDTO() {
    return new User(this.name, this.isVerified, this.email, this.password,
        this.contacts, this.resources, this.roles);
  }
}
