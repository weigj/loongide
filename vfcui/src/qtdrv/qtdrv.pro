#-------------------------------------------------
#
# Project created by QtCreator 2011-12-21T10:32:15
#
#-------------------------------------------------

TARGET = qtdrv
TEMPLATE = lib

DEFINES += QTDRV_LIBRARY

SOURCES += cdrv.cpp \
    qtsignal.cpp \
    qtevent.cpp \
    qtapp.cpp

HEADERS += cdrv.h\
        qtdrv_global.h \
    qtsignal.h \
    qtevent.h \
    qtapp.h

IDE_BUILD_TREE = $$PWD
IDE_APP_PATH = $$IDE_BUILD_TREE/../lib

DESTDIR = $$IDE_APP_PATH

symbian {
    MMP_RULES += EXPORTUNFROZEN
    TARGET.UID3 = 0xE1B1CE15
    TARGET.CAPABILITY = 
    TARGET.EPOCALLOWDLLDATA = 1
    addFiles.sources = qtdrv.dll
    addFiles.path = !:/sys/bin
    DEPLOYMENT += addFiles
}

unix:!symbian {
    maemo5 {
        target.path = /opt/usr/lib
    } else {
        target.path = /usr/lib
    }
    INSTALLS += target
}
