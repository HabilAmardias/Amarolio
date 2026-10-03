import { useEffect, useMemo } from "react";
import { ThemeProvider as MuiThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import { useAtom } from "jotai";
import { themeModeAtom } from "../store/atoms";
import { lightTheme, darkTheme, useSystemTheme } from "./theme";

export default function ThemeProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const systemMode = useSystemTheme();
  const [colorMode] = useAtom(themeModeAtom);
  const mode = colorMode === "system" ? systemMode : colorMode;

  const theme = useMemo(
    () => (mode === "dark" ? darkTheme : lightTheme),
    [mode]
  );

  useEffect(() => {
    document.documentElement.setAttribute("data-color-scheme", mode);
  }, [mode]);

  return (
    <MuiThemeProvider theme={theme}>
      <CssBaseline />
      {children}
    </MuiThemeProvider>
  );
}
