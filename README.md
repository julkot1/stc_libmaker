# STC libmaker file structure

---

## Table of Contents
- [Module Overview](#module-overview)
- [Head](#head)
  - [Includes](#includes)
  - [Types](#types)
    - [Type Methods](#type-methods)
    - [Type Match Rules](#type-match-rules)
- [Body](#body)
  - [Body Methods](#body-methods)

---

## Module Overview

- **Name**: Specifies the module name.
- **Includes**: Lists any header files included in the module.

```toml
name = "<module_name>"
includes = ["<header1.h>", "<header2.h>", ...]
```

---

## Head

The `[head]` section defines the module's structure and key data types.

### Includes

Specifies header files required for the module.

```toml
includes = ["<header1.h>", "<header2.h>", ...]
```

### Types

The `[head.types]` section defines specific types and their associated methods.

- **type_name**: The type's identifier within the module.
- **name**: Descriptive name for the type.
- **args**: List of argument types for functions associated with this type.
- **return**: The return type of the function.

```toml
[head.types]
type_name = "<type_identifier>"
name = "<type_name>"
args = ["<arg1_type>", "<arg2_type>", ...]
return = "<return_type>"
```

#### Type Methods

The `[head.types.method]` subsection defines method that matches function to arguments types.

- **name**: The name of the method.
- **stc**: Indicates if the method is available for rpn lang.
- **args**: Specifies argument types.
- **return**: The return type of the method.
- **code**: A list of code statements representing the method's logic.

```toml
[head.types.method]
name = "<method_name>"
stc = <true|false>
args = ["<arg1_type>", "<arg2_type>", ...]
return = "<return_type>"
code = ["<code_line1>", "<code_line2>", ...]
```

#### Type Match Rules

Each `[[head.types.match]]` table defines rules to match arguments for the specified type.

- **argA**: Type of the first argument in a match.
- **argB**: Type of the second argument in a match.
- **function**: Specifies the function to call when this match rule applies.

```toml
[[head.types.match]]
argA = "<argA_type>"
argB = "<argB_type>"
function = "<matching_function>"
```

---

## Body

The `[body]` section defines the main methods used by the module.

### Body Methods

Each `[[body.method]]` table specifies a method available in the module.

- **name**: The name of the method.
- **args**: List of argument types for the method.
- **return**: The method's return type.
- **stc**: Indicates if the method is available for rpn lang.
- **code**: Contains the list of code statements that define the method.

```toml
[[body.method]]
name = "<method_name>"
args = ["<arg1_type>", "<arg2_type>", ...]
return = "<return_type>"
stc = <true|false>
code = ["<code_line1>", "<code_line2>", ...]
```

---
