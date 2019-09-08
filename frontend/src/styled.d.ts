import "styled-components";

declare module "styled-components" {
  export interface DefaultTheme {
    primary: string;
    main: string;
    bg: string;
    text: string;
    border: string;
  }
}
