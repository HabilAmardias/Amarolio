import { Box, Container, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, TablePagination } from '@mui/material';
import { useAtom } from 'jotai';
import { urlHistoryAtom } from '../../models/url.model';
import { useState } from 'react';

export function DashboardPage() {
  const [urlHistory] = useAtom(urlHistoryAtom);
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(5);

  const handleChangePage = (
    event: React.MouseEvent<HTMLButtonElement> | null,
    newPage: number,
  ) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  return (
    <Container maxWidth="md">
      <Box sx={{ my: 4 }}>
        <Typography 
          variant="h3" 
          component="h1" 
          gutterBottom
          sx={{
            color: '#c25e00',
            fontWeight: 700,
            mb: 4,
          }}
        >
          Dashboard
        </Typography>

        <Typography 
          variant="h5" 
          sx={{ 
            mt: 4, 
            mb: 2,
            color: '#8b4513',
            fontWeight: 600,
          }}
        >
          Your Shortened URLs
        </Typography>

        {urlHistory.length === 0 ? (
          <Typography 
            sx={{ 
              color: 'text.secondary',
              fontSize: '1.1rem',
            }}
          >
            No URLs shortened yet.
          </Typography>
        ) : (
          <TableContainer component={Paper}>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell sx={{ color: '#c25e00', fontWeight: 700 }}>Short URL</TableCell>
                  <TableCell sx={{ color: '#c25e00', fontWeight: 700 }}>Original URL</TableCell>
                  <TableCell sx={{ color: '#c25e00', fontWeight: 700 }}>Expires At</TableCell>
                  <TableCell sx={{ color: '#c25e00', fontWeight: 700 }}>Created At</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {urlHistory
                  .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                  .map((url) => (
                  <TableRow 
                    key={url.shortUrl}
                    sx={{
                      '&:hover': {
                        background: '#faf6f0',
                      },
                    }}
                  >
                    <TableCell>
                      <a 
                        href={url.shortUrl} 
                        target="_blank" 
                        rel="noopener"
                        style={{
                          color: '#c25e00',
                          textDecoration: 'none',
                          '&:hover': {
                            textDecoration: 'underline',
                          },
                        }}
                      >
                        {url.shortUrl}
                      </a>
                    </TableCell>
                    <TableCell sx={{ maxWidth: 200, overflow: 'hidden', textOverflow: 'ellipsis', color: 'text.secondary' }}>
                      {url.originalUrl}
                    </TableCell>
                    <TableCell sx={{ color: '#8b4513' }}>
                      {url.expiresAt
                        ? new Date(url.expiresAt).toLocaleDateString()
                        : 'Never'}
                    </TableCell>
                    <TableCell sx={{ color: 'text.secondary' }}>
                      {new Date(url.createdAt).toLocaleDateString()}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
            <TablePagination
              rowsPerPageOptions={[5, 10, 25]}
              component="div"
              count={urlHistory.length}
              rowsPerPage={rowsPerPage}
              page={page}
              onPageChange={handleChangePage}
              onRowsPerPageChange={handleChangeRowsPerPage}
              sx={{
                borderTop: '1px solid #e8dcc8',
                color: '#5d4037',
                '& .MuiTablePagination-selectLabel, & .MuiTablePagination-displayedRows': {
                  color: '#5d4037',
                },
                '& .MuiIconButton-root': {
                  color: '#c25e00',
                  '&:hover': {
                    backgroundColor: 'rgba(194, 94, 0, 0.08)',
                  },
                },
              }}
            />
          </TableContainer>
        )}
      </Box>
    </Container>
  );
}
