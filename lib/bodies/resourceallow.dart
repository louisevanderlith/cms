import 'dart:html';

class ResourceAllow {
  LabelElement label;
  CheckboxInputElement checkbox;

  bool _loaded;

  ResourceAllow(HtmlElement item) {
    var li = item as LIElement;
    label = li.children[0] as LabelElement;
    checkbox = label.children[0] as CheckboxInputElement;

    _loaded = label != null && checkbox != null;
  }

  bool get loaded {
    return _loaded;
  }

  String get name {
    return label.innerText.replaceFirst(new RegExp(r'\n'), '').trim();
  }

  bool get allowed {
    return checkbox.checked;
  }
}
