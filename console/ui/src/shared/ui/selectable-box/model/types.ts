import { SxProps } from '@mui/material';
import { ReactNode } from 'react';

export interface ClusterFormSelectableBoxProps extends ReactNode {
  children?: ReactNode;
  isActive?: boolean;
  sx?: SxProps;
}
