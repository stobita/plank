import "styled-components";

declare module "styled-components" {
  export interface DefaultTheme {
    primary: string;
    primaryInner: string;
    main: string;
    bg: string;
    solid: string;
    weak: string;
    border: string;
    danger: string;
    dangerInner: string;
  }
}
