#ifndef QTSIGNAL_H
#define QTSIGNAL_H

#include <QObject>

class QtSignal : public QObject
{
    Q_OBJECT
public:
    explicit QtSignal(QObject *parent,void *func);
    virtual ~QtSignal();
public slots:
    void call();
    void call(bool);
    void call(int);
    void call(const QString&);
    void call(QObject*);
public:
    void *m_func;
};

#endif // QTSIGNAL_H
