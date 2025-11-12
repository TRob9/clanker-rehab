// Simple auto-close brackets for CodeMirror
(function(mod) {
  mod(CodeMirror);
})(function(CodeMirror) {

  var pairs = "()[]{}''\"\"";
  var explodeChars = "[]{}";

  CodeMirror.defineOption("autoCloseBrackets", false, function(cm, val, old) {
    if (old && old != CodeMirror.Init) {
      cm.removeKeyMap("autoCloseBrackets");
    }
    if (val) {
      var map = {name: "autoCloseBrackets"};

      // Add handlers for each opening bracket
      for (var i = 0; i < pairs.length; i += 2) {
        (function(left, right) {
          map["'" + left + "'"] = function(cm) {
            if (cm.somethingSelected()) {
              var selections = cm.getSelections();
              var newSelections = [];
              for (var i = 0; i < selections.length; i++) {
                newSelections.push(left + selections[i] + right);
              }
              cm.replaceSelections(newSelections, "around");
              return;
            }

            var cur = cm.getCursor();
            var line = cm.getLine(cur.line);
            var next = line.charAt(cur.ch);

            // Don't auto-close if next char is alphanumeric
            if (next && /\w/.test(next)) {
              return CodeMirror.Pass;
            }

            cm.replaceRange(left + right, cur);
            cm.setCursor({line: cur.line, ch: cur.ch + 1});
          };
        })(pairs.charAt(i), pairs.charAt(i + 1));
      }

      // Backspace handler
      map["Backspace"] = function(cm) {
        if (cm.somethingSelected()) return CodeMirror.Pass;

        var cur = cm.getCursor();
        var line = cm.getLine(cur.line);
        var prev = line.charAt(cur.ch - 1);
        var next = line.charAt(cur.ch);

        // Check if we're between a bracket pair
        for (var i = 0; i < pairs.length; i += 2) {
          if (pairs.charAt(i) === prev && pairs.charAt(i + 1) === next) {
            cm.replaceRange("",
              {line: cur.line, ch: cur.ch - 1},
              {line: cur.line, ch: cur.ch + 1});
            return;
          }
        }
        return CodeMirror.Pass;
      };

      // Enter handler for "exploding" brackets
      map["Enter"] = function(cm) {
        if (cm.somethingSelected()) return CodeMirror.Pass;

        var cur = cm.getCursor();
        var line = cm.getLine(cur.line);
        var prev = line.charAt(cur.ch - 1);
        var next = line.charAt(cur.ch);

        // Check if we're between explodable brackets
        var shouldExplode = false;
        for (var i = 0; i < explodeChars.length; i += 2) {
          if (explodeChars.charAt(i) === prev && explodeChars.charAt(i + 1) === next) {
            shouldExplode = true;
            break;
          }
        }

        if (shouldExplode) {
          // Get current line's indentation
          var currentIndent = line.substring(0, cur.ch - 1).match(/^\s*/)[0];

          // Determine indent character (tab or spaces)
          var indentUnit = cm.getOption("indentWithTabs") ? "\t" : " ".repeat(cm.getOption("indentUnit"));

          // Insert newlines with proper indentation
          var innerIndent = currentIndent + indentUnit;
          var outerIndent = currentIndent;

          cm.replaceSelection("\n" + innerIndent + "\n" + outerIndent, "end");

          // Move cursor to end of inner indentation
          var newCur = cm.getCursor();
          cm.setCursor({line: newCur.line - 1, ch: innerIndent.length});
          return;
        }

        return CodeMirror.Pass;
      };

      cm.addKeyMap(map);
    }
  });
});
