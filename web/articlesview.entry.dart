import 'package:CMS.APP/articleform.dart';
import 'package:mango_ui/keys.dart';

void main() {
  print('Article Edit');
  new ArticleForm(
      "#frmBlogCreate",
      getObjKey(),
      "#txtTitle",
      "#txtIntro",
      "#cboCategories",
      "#txtContent",
      "#uplHeaderImg",
      "#hdnAuthor",
      "#chkPublic",
      "#btnSubmit");
}
