import "styled-components";

declare module "styled-components" {
  export interface DefaultTheme {
    primary: string;
    bg: string;
    text: string;
    border: string;
  }
}
