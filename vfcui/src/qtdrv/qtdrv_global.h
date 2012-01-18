#ifndef QTDRV_GLOBAL_H
#define QTDRV_GLOBAL_H

#include <QtCore/qglobal.h>

#if defined(QTDRV_LIBRARY)
#  define QTDRVSHARED_EXPORT Q_DECL_EXPORT
#else
#  define QTDRVSHARED_EXPORT Q_DECL_IMPORT
#endif

#endif // QTDRV_GLOBAL_H
