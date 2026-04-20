import { createTheme } from "@mui/material/styles";

export const theme = createTheme({
  palette: {
    mode: "light",
    primary: { main: "#c25e00" }, // burnt orange
    secondary: { main: "#8b4513" }, // saddle brown
    background: { default: "#faf6f0", paper: "#ffffff" },
    text: { primary: "#3e2723", secondary: "#5d4037" },
  },
  typography: {
    fontFamily: "'Merriweather', 'Georgia', serif",
    button: { textTransform: "none", fontWeight: 600, letterSpacing: "0.5px" },
    h1: { fontWeight: 700, letterSpacing: "0.5px" },
    h2: { fontWeight: 700, letterSpacing: "0.5px" },
    h3: { fontWeight: 600, letterSpacing: "0.3px" },
    h4: { fontWeight: 600, letterSpacing: "0.3px" },
    h5: { fontWeight: 600 },
    h6: { fontWeight: 600 },
  },
  shape: { borderRadius: 8 },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          background: "linear-gradient(135deg, #c25e00 0%, #d2691e 100%)",
          color: "#ffffff",
          boxShadow: "0 2px 8px rgba(194, 94, 0, 0.3)",
          transition: "all 0.3s ease",
          "&:hover": {
            boxShadow: "0 4px 12px rgba(194, 94, 0, 0.4)",
            transform: "translateY(-1px)",
          },
        },
      },
    },
    MuiAppBar: {
      styleOverrides: {
        root: {
          background: "linear-gradient(90deg, #fdf6f0 0%, #f5e6d3 100%)",
          borderBottom: "1px solid #e0d5c1",
          boxShadow: "0 2px 8px rgba(139, 69, 19, 0.1)",
        },
      },
    },
    MuiTextField: {
      defaultProps: { size: "small", variant: "outlined" },
      styleOverrides: {
        root: {
          "& .MuiOutlinedInput-root": {
            borderColor: "#d4c5b0",
            "&:hover .MuiOutlinedInput-notchedOutline": {
              borderColor: "#c25e00",
            },
            "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
              borderColor: "#c25e00",
            },
          },
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          background: "#ffffff",
          border: "1px solid #e8dcc8",
          boxShadow: "0 2px 12px rgba(139, 69, 19, 0.08)",
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          background: "#ffffff",
          border: "1px solid #e8dcc8",
        },
      },
    },
    MuiAlert: {
      styleOverrides: {
        root: {
          border: "1px solid",
        },
      },
    },
  },
});
